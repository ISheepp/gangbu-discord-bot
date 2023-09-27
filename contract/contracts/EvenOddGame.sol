// SPDX-License-Identifier: MIT
// An example of a consumer contract that relies on a subscription for funding.
pragma solidity ^0.8.7;

import "@chainlink/contracts/src/v0.8/interfaces/VRFCoordinatorV2Interface.sol";
import "@chainlink/contracts/src/v0.8/VRFConsumerBaseV2.sol";
import "@chainlink/contracts/src/v0.8/ConfirmedOwner.sol";

/// @title even odd game
/// @author ISheep
/// @notice none
contract EvenOddGame is VRFConsumerBaseV2, ConfirmedOwner {
    enum Choice {
        Even,
        Odd
    }

    enum GameState {
        InProgress,
        Finished
    }

    error CallFailed();

    event RequestSent(uint256 requestId, address caller);
    event RequestFulfilled(uint256 requestId, uint256[] randomWords);
    event BetResult(
        uint256 requestId,
        uint256 amount,
        address callerAddress,
        bool gameResult
    );
    struct Bet {
        bool fulfilled; // whether the request has been successfully fulfilled
        bool exists; // whether a requestId exists
        uint256 randomNum; // random num
        address payable callerAddress;
        uint256 amount;
        Choice choice; // 0 even 1 odd
        GameState gameStatus;
        uint256 gameDoneTimestamp;
    }
    mapping(uint256 => Bet) public s_requests; /* requestId --> bet */

    mapping(address => uint256[])
        public addr_requests; /* address --> requestIds[] */

    VRFCoordinatorV2Interface COORDINATOR;

    // Your subscription ID.
    uint64 s_subscriptionId;

    // past requests Id.
    uint256[] public requestIds;
    uint256 public lastRequestId;

    // chainlink configuration
    // see https://docs.chain.link/docs/vrf/v2/subscription/supported-networks/#configurations
    bytes32 keyHash =
        0x79d3d8832d904592c0bf9818b621522c988bb8b0c05cdc3b15aea1b6e8db0c15;
    uint32 callbackGasLimit = 100000;
    uint16 requestConfirmations = 3;
    uint32 numWords = 1;

    constructor(
        uint64 subscriptionId
    )
        payable
        VRFConsumerBaseV2(0x2Ca8E0C643bDe4C2E08ab1fA0da3401AdAD7734D)
        ConfirmedOwner(msg.sender)
    {
        COORDINATOR = VRFCoordinatorV2Interface(
            0x2Ca8E0C643bDe4C2E08ab1fA0da3401AdAD7734D
        );
        s_subscriptionId = subscriptionId;
    }

    /// @dev request random
    /// @param _choice player's choice
    /// @return requestId requestId from VRF
    function requestRandomWords(
        uint8 _choice
    ) external payable returns (uint256 requestId) {
        require(msg.value > 0, "amount must be greater than 0");
        // Will revert if subscription is not set and funded.
        requestId = COORDINATOR.requestRandomWords(
            keyHash,
            s_subscriptionId,
            requestConfirmations,
            callbackGasLimit,
            numWords
        );
        Choice finalChoice = _choice == 0 ? Choice.Even : Choice.Odd;
        Bet memory newBet = Bet({
            randomNum: 0,
            callerAddress: payable(msg.sender),
            exists: true,
            fulfilled: false,
            amount: msg.value,
            gameDoneTimestamp: 0,
            choice: finalChoice,
            gameStatus: GameState.InProgress
        });
        addr_requests[msg.sender].push(requestId);
        s_requests[requestId] = newBet;
        requestIds.push(requestId);
        lastRequestId = requestId;
        emit RequestSent(requestId, msg.sender);
        return requestId;
    }

    /// @dev main game logic
    /// @param _requestId requestId
    /// @return result game result
    function mainBet(uint256 _requestId) external returns (bool result) {
        require(
            s_requests[_requestId].callerAddress == msg.sender,
            "game must be caller"
        );
        require(s_requests[_requestId].exists, "request not found");
        require(
            s_requests[_requestId].fulfilled,
            "random num is not generated"
        );
        require(
            s_requests[_requestId].gameStatus == GameState.InProgress,
            "this game request already done"
        );
        Bet storage bet = s_requests[_requestId];
        bool isGuessCorrect = (bet.randomNum % 2 == 0 &&
            bet.choice == Choice.Even) ||
            (bet.randomNum % 2 != 0 && bet.choice == Choice.Odd);
        if (isGuessCorrect) {
            uint256 winnings = bet.amount * 2; // 1:1 payout
            sendReward(bet.callerAddress, winnings);
        }
        bet.gameStatus = GameState.Finished;
        emit BetResult(_requestId, bet.amount, bet.callerAddress, isGuessCorrect);
        return isGuessCorrect;
    }

    function fulfillRandomWords(
        uint256 _requestId,
        uint256[] memory _randomWords
    ) internal override {
        require(s_requests[_requestId].exists, "request not found");
        s_requests[_requestId].fulfilled = true;
        s_requests[_requestId].randomNum = _randomWords[0];
        s_requests[_requestId].gameDoneTimestamp = block.timestamp;
        emit RequestFulfilled(_requestId, _randomWords);
    }

    /// @dev get request status by requestId
    /// @param _requestId requestId
    /// @return fulfilled is random num generate complete
    /// @return callerAddress caller EOA address
    /// @return randomNum random num
    function getRequestStatus(
        uint256 _requestId
    )
        external
        view
        returns (bool fulfilled, address callerAddress, uint256 randomNum)
    {
        require(s_requests[_requestId].exists, "request not found");
        Bet memory thisBet = s_requests[_requestId];
        return (thisBet.fulfilled, thisBet.callerAddress, thisBet.randomNum);
    }

    function withdraw() external onlyOwner {
        uint256 contractBalance = address(this).balance;
        require(contractBalance > 0, "Contract has no balance to withdraw");
        payable(owner()).transfer(contractBalance);
    }

    /// sendETH util
    /// @param _to to address
    /// @param amount ETH amount wei
    function sendReward(address payable _to, uint256 amount) internal {
        // 处理下call的返回值，如果失败，revert交易并发送error
        require(
            address(this).balance > amount,
            "contract has insufficient ETH"
        );
        (bool success, ) = _to.call{value: amount}("");
        if (!success) {
            revert CallFailed();
        }
    }

    receive() external payable {}
}

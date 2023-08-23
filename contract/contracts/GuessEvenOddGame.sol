// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@chainlink/contracts/src/v0.8/VRFConsumerBase.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract GuessEvenOddGame is VRFConsumerBase, Ownable {
    bytes32 internal keyHash;
    uint256 internal fee;

    enum Guess { Even, Odd }
    enum GameState { InProgress, Finished }

    struct Bet {
        address player;
        Guess guess;
        uint256 amount;
        GameState state;
    }

    mapping(bytes32 => Bet) public bets;
    mapping(address => uint256) public balances;

    event BetPlaced(address indexed player, Guess indexed guess, uint256 amount, bytes32 requestId);
    event BetResult(address indexed player, Guess indexed guess, bool indexed win, uint256 amount);

    constructor(address vrfCoordinator, address linkToken, bytes32 _keyHash, uint256 _fee)
        VRFConsumerBase(vrfCoordinator, linkToken)
    {
        keyHash = _keyHash;
        fee = _fee;
    }

    function placeBet(Guess guess) external payable {
        require(msg.value > 0, "Amount must be greater than 0");
        require(guess == Guess.Even || guess == Guess.Odd, "Invalid guess");

        bytes32 requestId = requestRandomness(keyHash, fee);
        bets[requestId] = Bet(msg.sender, guess, msg.value, GameState.InProgress);

        emit BetPlaced(msg.sender, guess, msg.value, requestId);
    }

    function fulfillRandomness(bytes32 requestId, uint256 randomness) internal override {
        Bet storage bet = bets[requestId];
        require(bet.state == GameState.InProgress, "Bet not in progress");

        bool isEven = randomness % 2 == 0;
        bool isGuessCorrect = (isEven && bet.guess == Guess.Even) || (!isEven && bet.guess == Guess.Odd);

        if (isGuessCorrect) {
            uint256 winnings = bet.amount * 2; // 1:1 payout
            balances[bet.player] += winnings;
        }

        bet.state = GameState.Finished;
        emit BetResult(bet.player, bet.guess, isGuessCorrect, winnings);
    }

    function withdraw() external {
        uint256 balance = balances[msg.sender];
        require(balance > 0, "No balance to withdraw");

        balances[msg.sender] = 0;
        payable(msg.sender).transfer(balance);
    }

    function setKeyHash(bytes32 _keyHash) external onlyOwner {
        keyHash = _keyHash;
    }

    function setFee(uint256 _fee) external onlyOwner {
        fee = _fee;
    }
}

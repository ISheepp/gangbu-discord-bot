const {
  time,
  loadFixture,
} = require("@nomicfoundation/hardhat-toolbox/network-helpers");
const { anyValue } = require("@nomicfoundation/hardhat-chai-matchers/withArgs");
const { expect } = require("chai");

describe("测试Show", function () {
  // 部署合约
  async function deployShowContract() {
    const ONE_GWEI = 1_000_000_000;
    const lockedAmount = ONE_GWEI;

    const [owner, otherAccount] = await ethers.getSigners();

    const Show = await ethers.getContractFactory("Show");
    const show = await Show.deploy();

    return { show, lockedAmount, owner, otherAccount };
  }

  describe("部署测试", function() {
    it("加上一个数应该为1", async function () {
      console.log("hihihi");
      const { show } = await loadFixture(deployShowContract);

      await show.add()
      expect(await show.num()).to.equal(1);
    });
  })

});

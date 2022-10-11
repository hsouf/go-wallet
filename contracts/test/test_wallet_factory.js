const Factory = artifacts.require("WalletFactory");
const Wallet = artifacts.require("Wallet");
const ethers = require("ethers");
const {
  expectEvent, // Assertions for emitted events
  expectRevert, // Assertions for transactions that should fail
} = require("@openzeppelin/test-helpers");

require("chai").use(require("chai-as-promised")).should();

contract("WalletFactory", (accounts) => {
  let factory;
  let factoryGov;
  let walletAdmin;
  let mintCost = 0.003 * 10 ** 18;
  let provider;

  before(async () => {
    [factoryGov, walletAdmin] = accounts.slice(0, 2);
    provider = new ethers.providers.JsonRpcProvider("http://127.0.0.1:9999");
    factory = await Factory.new(factoryGov);
  });

  describe("Factory mints and intializes wallet correctly ", async () => {
    it("mints wallet correctly", async () => {
      const mintTx = await factory.mintWallet(walletAdmin, {
        value: mintCost,
      });

      expectEvent(mintTx, "WalletMinted");

      //get event args
      const mintEvent = mintTx.logs[0].args;

      //get address of newly created wallet
      const walletAddress = mintEvent.wallet;

      //check that admin was initialized correctly
      expect(walletAdmin).to.equal(
        ethers.utils.getAddress(await provider.getStorageAt(walletAddress, 0))
      );
    });

    it("should fail if wrong mint cost value is provided", async () => {
      await expectRevert(
        factory.mintWallet(walletAdmin, {
          value: mintCost - 1,
        }),
        "Factory: insufficient value"
      );
    });
  });
});

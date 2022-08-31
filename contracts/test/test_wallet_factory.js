const Factory = artifacts.require("WalletFactory");
const Wallet = artifacts.require("Wallet");

const {
  BN, // Big Number support
  constants, // Common constants, like the zero address and largest integers
  expectEvent, // Assertions for emitted events
  expectRevert, // Assertions for transactions that should fail
} = require("@openzeppelin/test-helpers");

require("chai").use(require("chai-as-promised")).should();

contract("WalletFactory", (accounts) => {
  let factory;
  let factoryGov;
  let walletAdmin;
  let mintCost = 0.003 * 10 ** 18;

  before(async () => {
    [factoryGov, walletAdmin] = accounts.slice(0, 2);
    console.log(factoryGov);
    console.log(walletAdmin);
    factory = await Factory.new(factoryGov);
  });

  describe("Wallets minting ", async () => {
    it("mints wallet correctly", async () => {
      const mintTx = await factory.mintWallet(walletAdmin, {
        value: mintCost,
      });

      expectEvent(mintTx, "WalletMinted");

      //get event args
      const mintEvent = mintTx.logs[0].args;
      const admin = mintEvent.admin;
      const walletAddress = mintEvent.wallet;

      const wallet = await Wallet.at(walletAddress);

      //check that admin was initialized correctly
      //expect(walletAdmin).to.equal(await wallet.admin());
    });

    it("should fail if wrong mint cost value is provided", async () => {
      expectRevert(
        await factory.mintWallet(walletAdmin, {
          value: mintCost - 1,
        }),
        "Factory: insufficient value"
      );
    });
  });
});

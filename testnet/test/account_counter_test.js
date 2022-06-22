const AccountCounterTest = artifacts.require("AccountCounter");

/*
 * uncomment accounts to access the test accounts made available by the
 * Ethereum client
 * See docs: https://www.trufflesuite.com/docs/truffle/testing/writing-tests-in-javascript
 */
contract("AccountCounterTest", function (accounts) {
  it("Increment the account value", async function () {
    const instance = await AccountCounterTest.deployed();
    await instance.incrementCount.call({from: accounts[0]});
    const count_value = (await instance.getCount.call(accounts[0])).toNumber();
    return assert.equal(1, count_value);
  });
});

const AccountCounter = artifacts.require("AccountCounter");

module.exports = function (deployer) {
    deployer.deploy(AccountCounter);
};

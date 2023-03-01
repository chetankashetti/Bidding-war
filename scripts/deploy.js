const hre = require("hardhat");

async function main() {
  // We get the contract to deploy
  const BidWar = await hre.ethers.getContractFactory("BidWar");
  const bidwar = await BidWar.deploy(3600, 600, 5);

  await bidwar.deployed();

  console.log("Bidwar deployed to:", bidwar.address);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });

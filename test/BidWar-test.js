const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("BidWar", function() {
    it("should throw error if bid bid amount is lower", async function () {
        const BidWar = await ethers.getContractFactory("BidWar");
        bidWar = await BidWar.deploy(3600, 600, 5);
        const Energi = await ethers.getContractFactory("EnergiToken");
        const energiToken = await Energi.deploy(10000);


        await expect(
            bidWar.newBid(0, energiToken.address)
        ).to.be.revertedWith("LowerBidAmount");
    })

    it("should throw incorrect allowance error if there's low balance", async function () {
        const [chetan, a1, a2] = await ethers.getSigners();
        const BidWar = await ethers.getContractFactory("BidWar");
        bidWar = await BidWar.deploy(3600, 600, 5);

        const Energi = await ethers.getContractFactory("EnergiToken");
        const energiToken = await Energi.deploy(10000);

        await expect(
            bidWar.newBid(100, energiToken.address)
        ).to.be.revertedWith("IncorrectAllowance");
    })

    it("should emit log", async function () {
        const [chetan, a1, a2] = await ethers.getSigners();
        const BidWar = await ethers.getContractFactory("BidWar");
        const bidWar = await BidWar.deploy(3600, 600, 5);

        const Energi = await ethers.getContractFactory("EnergiToken");
        const energiToken = await Energi.deploy(10000);
        
        await energiToken.approve(bidWar.address, energiToken.totalSupply())

        await expect(bidWar.newBid(10, energiToken.address))
        .to.emit(bidWar, 'NewBid')
    })

    it("should transfer commission to owner", async function () {
        const [chetan, a1, a2] = await ethers.getSigners();
        const BidWar = await ethers.getContractFactory("BidWar");
        const bidWar = await BidWar.deploy(3600, 600, 5);

        const Energi = await ethers.getContractFactory("EnergiToken");
        const energiToken = await Energi.deploy(10000);
        
        await energiToken.approve(bidWar.address, energiToken.totalSupply())

        await bidWar.newBid(100, energiToken.address)
        balance = await energiToken.balanceOf(chetan.address);
        expect(balance).to.equals(9905)
    })

    it("should transfer bid amount to contract", async function () {
        const [chetan, a1, a2] = await ethers.getSigners();
        const BidWar = await ethers.getContractFactory("BidWar");
        const bidWar = await BidWar.deploy(3600, 600, 5);

        const Energi = await ethers.getContractFactory("EnergiToken");
        const energiToken = await Energi.deploy(10000);
        
        await energiToken.approve(bidWar.address, energiToken.totalSupply())
        await bidWar.newBid(10, energiToken.address)

        balance = await energiToken.balanceOf(bidWar.address);
        expect(balance).to.equals(10)
    })

    it("multiple bids should work fine", async function () {
        const [chetan, a1, a2] = await ethers.getSigners();
        const BidWar = await ethers.getContractFactory("BidWar");
        const bidWar = await BidWar.deploy(3600, 600, 5);
        const Energi = await ethers.getContractFactory("EnergiToken");
        const energiToken = await Energi.deploy(10000);
        
        await energiToken.approve(bidWar.address, energiToken.totalSupply())
        await bidWar.newBid(10, energiToken.address)
        balance = await energiToken.balanceOf(bidWar.address);
        expect(balance).to.equals(10)

        energiToken.transfer(a1.address, 100)
        await energiToken.connect(a1).approve(bidWar.address, 100)
        await bidWar.connect(a1).newBid(12, energiToken.address)
        balance = await energiToken.balanceOf(bidWar.address);
        expect(balance).to.equals(22)
    })

    it("should revert for finalise before timeout", async function() {
        const [chetan, a1, a2] = await ethers.getSigners();
        const BidWar = await ethers.getContractFactory("BidWar");
        const bidWar = await BidWar.deploy(3600, 600, 5);
        const Energi = await ethers.getContractFactory("EnergiToken");
        const energiToken = await Energi.deploy(10000);
        
        await energiToken.approve(bidWar.address, energiToken.totalSupply())
        await bidWar.newBid(10, energiToken.address)

        await expect(
            bidWar.finalise(energiToken.address)
        ).to.be.revertedWith("BidTimeLeft");
    })

    it("should finalise after timeout and transfer bid amount to winner", async function() {
        const [chetan, a1, a2] = await ethers.getSigners();
        const BidWar = await ethers.getContractFactory("BidWar");
        const bidWar = await BidWar.deploy(1, 1, 5);
        const Energi = await ethers.getContractFactory("EnergiToken");
        const energiToken = await Energi.deploy(10000);
        
        await energiToken.approve(bidWar.address, energiToken.totalSupply())
        await bidWar.newBid(20, energiToken.address)

        energiToken.transfer(a1.address, 100)
        await energiToken.connect(a1).approve(bidWar.address, 100)
        await bidWar.connect(a1).newBid(30, energiToken.address)
        balance = await energiToken.balanceOf(a1.address);
        expect(balance).to.equals(70)

        await bidWar.newBid(40, energiToken.address)
        await bidWar.connect(a1).newBid(50, energiToken.address)

        await bidWar.finalise(energiToken.address)
        balance = await energiToken.balanceOf(a1.address);
        expect(balance).to.equals(154)
    })

    it("should emit finalise log", async function() {
        const [chetan, a1, a2] = await ethers.getSigners();
        const BidWar = await ethers.getContractFactory("BidWar");
        const bidWar = await BidWar.deploy(1, 1, 5);
        const Energi = await ethers.getContractFactory("EnergiToken");
        const energiToken = await Energi.deploy(10000);
        
        await energiToken.approve(bidWar.address, energiToken.totalSupply())
        await bidWar.newBid(100, energiToken.address)

        energiToken.transfer(a1.address, 1000)
        await energiToken.connect(a1).approve(bidWar.address, 1000)
        await bidWar.connect(a1).newBid(200, energiToken.address)

        await expect(bidWar.finalise(energiToken.address))
        .to.emit(bidWar, 'Timeout')
        .withArgs(a1.address, 285)
    })


});
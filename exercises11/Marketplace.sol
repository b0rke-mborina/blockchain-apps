// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0 <0.9.0;

contract Marketplace {
    // structure for items being sold
    struct Item {
        string name;
        string description;
        uint price;
        address seller;
    }

    // mapping to keep track of the items being sold
    mapping(uint => Item) public items;

    // keeps count of number of items on the marketplace
    uint public itemCount = 0;

    // modifier to ensure that only the seller of an item can remove it from the marketplace
    modifier onlySeller(uint itemId) {
        require(items[itemId].seller == msg.sender, "Only the seller can remove an item.");
        _;
    }

    // modifier to ensure that only authorized buyers can purchase items from the marketplace
    modifier onlyBuyer(uint itemId) {
        require(msg.sender != items[itemId].seller, "The seller cannot buy their own item.");
        _;
    }

    // modifier to ensure that the correct amount is sent
    modifier correctAmount(uint itemId) {
        require(items[itemId].price == msg.value, "Incorrect amount sent.");
        _;
    }

    // function for adding a new item to the marketplace
    function addItem(string memory name, string memory description, uint price) public {
        Item memory newItem = Item({
            name: name,
            description: description,
            price: price,
            seller: msg.sender
        });

        items[itemCount] = newItem;
        itemCount++;
    }

    // function for buying an item from the marketplace
    function buyItem(uint itemId) public payable onlyBuyer(itemId) correctAmount(itemId) {
        address payable seller = payable(items[itemId].seller);
        seller.transfer(msg.value);

        delete items[itemId];
        itemCount--;
    }

    // function for removing an item from the marketplace
    function removeItem(uint itemId) public onlySeller(itemId) {
        delete items[itemId];
        itemCount--;
    }

    // function for retrieving the details of an item
    function getItem(uint itemId) public view returns (string memory, string memory, uint, address) {
        Item memory item = items[itemId];
        return (item.name, item.description, item.price, item.seller);
    }

    //  function for retrieving the number of items currently for sale on the marketplace
    function getNumItems() public view returns (uint) {
        return itemCount;
    }
}
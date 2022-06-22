// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

contract AccountCounter {
  // Hesap ve sayaç bilgisini tut
  mapping(address => uint32) counts;

  // Sayaç bilgisini arttır
  function incrementCount() public {
    // Sayaç bilgisinin boyutu 32bit, 
    // eğer bundan yukarı çıkacaksak hata döndür
    if (counts[msg.sender] >= type(uint32).max)
      revert("Cannot increment anymore, max");

    // Sayaç değerini arttır
    counts[msg.sender] += 1;
  }

  // Sayaç bilgisini getir
  function getCount(address addr) public view returns(uint32) {
    return counts[addr];
  }
}

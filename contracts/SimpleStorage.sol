// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract SimpleStorage {
    // 定义一个映射，用于存储键值对
    mapping(uint256 => string) public dataMap;

    // 存储函数：接受一个整数作为键，字符串作为值
    function set(uint256 key, string memory value) public {
        dataMap[key] = value;
    }

    // 读取函数：通过给定的键返回相应的值
    function get(uint256 key) public view returns (string memory) {
        return dataMap[key];
    }
}
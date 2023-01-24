// SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.7.0 <0.9.0;

contract Contract {
    address public owner;
    Task[] tasks;

    struct Task {
        string content;
        bool status;
    }

    constructor() {
        // owner = msg.sender;
    }

    // modifier isOwner() {
    //     require(owner == msg.sender);
    //     _;
    // }

    function add(string memory _content) public {
        tasks.push(Task(_content, false));
    }

    function get(uint256 _id) public view returns (Task memory) {
        return tasks[_id];
    }

    function list() public view returns (Task[] memory) {
        return tasks;
    }

    function update(uint256 _id, string memory _content) public {
        tasks[_id].content = _content;
    }

    function toggle(uint _id) public {
      tasks[_id].status = !tasks[_id].status;
    }

    function remove(uint256 _id) public {
        for (uint256 i = _id; i < tasks.length - 1; i++) {
            tasks[i] = tasks[i + 1];
        }
        tasks.pop();
    }
}

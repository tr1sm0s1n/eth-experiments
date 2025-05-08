// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

contract Registry {
    mapping(string => mapping(uint => string)) properties;
    mapping(string => uint) propertyVersions;

    function addProperty(string[] memory landIds, string[] memory data) public {
        for (uint i = 0; i < landIds.length; i++) {
            uint newVersion = propertyVersions[landIds[i]] + 1;
            properties[landIds[i]][newVersion] = data[i];
            propertyVersions[landIds[i]] = newVersion;
        }
    }

    function getLatestProperty(
        string memory propertyId
    ) public view returns (string memory) {
        uint latestVersion = propertyVersions[propertyId];
        return properties[propertyId][latestVersion];
    }

    function getVersionProperty(
        string memory propertyId,
        uint version
    ) public view returns (string memory) {
        return properties[propertyId][version];
    }

    function getPropertyVersions(
        string memory propertyId
    ) public view returns (uint) {
        return propertyVersions[propertyId];
    }

    function getAllPropertyVersions(
        string memory propertyId
    ) public view returns (string[] memory) {
        uint versions = propertyVersions[propertyId];
        string[] memory allVersions = new string[](versions);
        for (uint i = 1; i <= versions; i++) {
            allVersions[i - 1] = properties[propertyId][i];
        }
        return allVersions;
    }
}

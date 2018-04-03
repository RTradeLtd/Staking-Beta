pragma solidity 0.4.21;

contract EthRouter {
	// used to send eth to multiple people

	event EthSent(address indexed _recipient, uint256 _amount);

	function sendEth(
		address[] _addrs,
		uint256	_ethPerAddress)
		public
		payable
		returns (bool)
	{
		require(msg.value >= (_addrs.length * _ethPerAddress));
		for (uint256 i = 0; i < _addrs.length; i++) {
			if (_addrs[i] != address(0)) {
				require(_addrs[i].send(_ethPerAddress));
				emit EthSent(_addrs[i], _ethPerAddress);
			}
		}
		return true;
	}


	function sendEth(
		address[] _addrs,
		uint256[] _eth)
		public
		payable
		returns (bool)
	{
		require(msg.value > 0);
		for (uint256 i = 0; i < _addrs.length; i++) {
			if (_addrs[i] != address(0) && _eth[i] > 0) {
				require(_addrs[i].send(_eth[i]));
				emit EthSent(_addrs[i], _eth[i]);
			}
		}
		return true;
	}

}
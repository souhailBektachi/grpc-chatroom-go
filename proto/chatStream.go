syntaxt ="proto3";

option go_package = "./proto";

package chatStream;

service ChatStream {
	rpc Chat(stream ChatMessage) returns (stream ChatMessage) {}
}

message ChatMessage {
	string user = 1;
	string message = 2;
}
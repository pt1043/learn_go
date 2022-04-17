package apis


service Chatservice{
rpc Register(Personalinformation) returns (Personalinformation)
rpc WatchPersons(null)returns (stream PersonalInformation){}
}
message null {}
message Personalinformation
string Name = 1
string Password = 2
string message = 3

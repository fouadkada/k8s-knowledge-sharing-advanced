// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var search_pb = require('./search_pb.js');

function serialize_proto_Movie(arg) {
  if (!(arg instanceof search_pb.Movie)) {
    throw new Error('Expected argument of type proto.Movie');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_Movie(buffer_arg) {
  return search_pb.Movie.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_Results(arg) {
  if (!(arg instanceof search_pb.Results)) {
    throw new Error('Expected argument of type proto.Results');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_Results(buffer_arg) {
  return search_pb.Results.deserializeBinary(new Uint8Array(buffer_arg));
}


var SearchService = exports.SearchService = {
  search: {
    path: '/proto.Search/Search',
    requestStream: false,
    responseStream: false,
    requestType: search_pb.Movie,
    responseType: search_pb.Results,
    requestSerialize: serialize_proto_Movie,
    requestDeserialize: deserialize_proto_Movie,
    responseSerialize: serialize_proto_Results,
    responseDeserialize: deserialize_proto_Results,
  },
};

exports.SearchClient = grpc.makeGenericClientConstructor(SearchService);

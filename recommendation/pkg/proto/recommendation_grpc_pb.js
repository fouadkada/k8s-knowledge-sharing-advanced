// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var recommendation_pb = require('./recommendation_pb.js');

function serialize_proto_Movie(arg) {
  if (!(arg instanceof recommendation_pb.Movie)) {
    throw new Error('Expected argument of type proto.Movie');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_Movie(buffer_arg) {
  return recommendation_pb.Movie.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_Results(arg) {
  if (!(arg instanceof recommendation_pb.Results)) {
    throw new Error('Expected argument of type proto.Results');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_Results(buffer_arg) {
  return recommendation_pb.Results.deserializeBinary(new Uint8Array(buffer_arg));
}


var RecommendationService = exports.RecommendationService = {
  getRecommendations: {
    path: '/proto.Recommendation/GetRecommendations',
    requestStream: false,
    responseStream: false,
    requestType: recommendation_pb.Movie,
    responseType: recommendation_pb.Results,
    requestSerialize: serialize_proto_Movie,
    requestDeserialize: deserialize_proto_Movie,
    responseSerialize: serialize_proto_Results,
    responseDeserialize: deserialize_proto_Results,
  },
};

exports.RecommendationClient = grpc.makeGenericClientConstructor(RecommendationService);

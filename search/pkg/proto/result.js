/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

goog.provide('proto.proto.Result');

goog.require('jspb.BinaryReader');
goog.require('jspb.BinaryWriter');
goog.require('jspb.Message');


/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.proto.Result = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.proto.Result, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.proto.Result.displayName = 'proto.proto.Result';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.proto.Result.prototype.toObject = function(opt_includeInstance) {
  return proto.proto.Result.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.proto.Result} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.Result.toObject = function(includeInstance, msg) {
  var f, obj = {
    voteaverage: +jspb.Message.getFieldWithDefault(msg, 1, 0.0),
    id: jspb.Message.getFieldWithDefault(msg, 2, 0),
    releasedate: jspb.Message.getFieldWithDefault(msg, 3, ""),
    backdroppath: jspb.Message.getFieldWithDefault(msg, 4, ""),
    title: jspb.Message.getFieldWithDefault(msg, 5, ""),
    posterpath: jspb.Message.getFieldWithDefault(msg, 6, ""),
    overview: jspb.Message.getFieldWithDefault(msg, 7, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.proto.Result}
 */
proto.proto.Result.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.proto.Result;
  return proto.proto.Result.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.proto.Result} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.proto.Result}
 */
proto.proto.Result.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readFloat());
      msg.setVoteaverage(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setReleasedate(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setBackdroppath(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setTitle(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setPosterpath(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.setOverview(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.proto.Result.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.proto.Result.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.proto.Result} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.Result.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getVoteaverage();
  if (f !== 0.0) {
    writer.writeFloat(
      1,
      f
    );
  }
  f = message.getId();
  if (f !== 0) {
    writer.writeInt64(
      2,
      f
    );
  }
  f = message.getReleasedate();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getBackdroppath();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getTitle();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getPosterpath();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
  f = message.getOverview();
  if (f.length > 0) {
    writer.writeString(
      7,
      f
    );
  }
};


/**
 * optional float VoteAverage = 1;
 * @return {number}
 */
proto.proto.Result.prototype.getVoteaverage = function() {
  return /** @type {number} */ (+jspb.Message.getFieldWithDefault(this, 1, 0.0));
};


/** @param {number} value */
proto.proto.Result.prototype.setVoteaverage = function(value) {
  jspb.Message.setField(this, 1, value);
};


/**
 * optional int64 ID = 2;
 * @return {number}
 */
proto.proto.Result.prototype.getId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {number} value */
proto.proto.Result.prototype.setId = function(value) {
  jspb.Message.setField(this, 2, value);
};


/**
 * optional string ReleaseDate = 3;
 * @return {string}
 */
proto.proto.Result.prototype.getReleasedate = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.proto.Result.prototype.setReleasedate = function(value) {
  jspb.Message.setField(this, 3, value);
};


/**
 * optional string BackdropPath = 4;
 * @return {string}
 */
proto.proto.Result.prototype.getBackdroppath = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.proto.Result.prototype.setBackdroppath = function(value) {
  jspb.Message.setField(this, 4, value);
};


/**
 * optional string Title = 5;
 * @return {string}
 */
proto.proto.Result.prototype.getTitle = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/** @param {string} value */
proto.proto.Result.prototype.setTitle = function(value) {
  jspb.Message.setField(this, 5, value);
};


/**
 * optional string PosterPath = 6;
 * @return {string}
 */
proto.proto.Result.prototype.getPosterpath = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/** @param {string} value */
proto.proto.Result.prototype.setPosterpath = function(value) {
  jspb.Message.setField(this, 6, value);
};


/**
 * optional string Overview = 7;
 * @return {string}
 */
proto.proto.Result.prototype.getOverview = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/** @param {string} value */
proto.proto.Result.prototype.setOverview = function(value) {
  jspb.Message.setField(this, 7, value);
};



import grpc from "@grpc/grpc-js";
import protoLoader from "@grpc/proto-loader";

var PROTO_PATH = __dirname + "/../../../protos/ml.proto";

/**
 * Retrieves a tokenizer client instance for the given URL.
 *
 * @param {string} [url] - The optional URL to get the client for.
 * @returns {Object} The client instance.
 */
export default function GetTokenizerClient(url) {
  if (url === undefined) {
    url = "localhost:50051";
  }
  let packageDefinition = protoLoader.loadSync(PROTO_PATH, {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true,
  });

  let token_proto = grpc.loadPackageDefinition(packageDefinition);

  let client = new token_proto.TokensServer(
    url,
    grpc.credentials.createInsecure(),
  );
  return client;
}

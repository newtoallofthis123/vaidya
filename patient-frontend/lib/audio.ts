import grpc, { GrpcObject } from "@grpc/grpc-js";
import protoLoader from "@grpc/proto-loader";

var PROTO_PATH = "@/lib/transcribe.proto";

export function GetTranscribeClient(url?: string) {
  console.log("URL", PROTO_PATH);
  if (url === undefined) {
    url = "localhost:50052";
  }
  let packageDefinition = protoLoader.loadSync(PROTO_PATH, {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true,
  });

  let transcribe_proto = grpc.loadPackageDefinition(packageDefinition);

  let client = new transcribe_proto.AudioService(
    url,
    grpc.credentials.createInsecure(),
    {},
  );
  return client;
}

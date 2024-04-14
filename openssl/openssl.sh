# Output files
# ca.key: Certificate Authority private key file (in real-life : this shouldn't be shared)
# ca.crt: Certificate Authority trust certificate (in real-life : share this with gRPC client)
# server.key: Server private key, password protected (in real-life : this shouldn't be shared)
# server.csr: Server certificate signing request (in real-life : share this with CA owner)
# server.crt: Server certificate to be "installed" on server (in real-life : this shouldn't be shared)
# server.pem: server.key, converted into a format required by gRPC (in real-life : this shouldn't be shared)
#
# Private files (not to be shared): ca.key, server.key, server.pem, server.crt
# Public files (to be shared): ca.crt (share to gRPC client), server.csr (share to CA)

SERVER_CN=localhost
MY_SUBJECT="/CN=${SERVER_CN}"

# 1: Generate Certificate Authority + Trust Certificate (ca.crt)
openssl genrsa -passout pass:1234 -des3 -out ca.key 4096
openssl req -passin pass:1234 -new -x509 -sha256 -days 365 -key ca.key -out ca.crt -subj "/CN=ca"

# 2: Generate server private key (server.key)
openssl genrsa -passout pass:1234 -des3 -out server.key 4096

# 3: Get a certificate signing request from the CA (server.csr)
openssl req -passin pass:1234 -new -key server.key -out server.csr -subj ${MY_SUBJECT}

# 4: Self-sign the certificate with the CA that just created (server.crt)
openssl x509 -req -extfile <(printf "subjectAltName=DNS:${SERVER_CN}") -passin pass:1234 -sha256 -days 365 -in server.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out server.crt 

# 5: Convert the server private key to .pem format that required by gRPC (server.pem)
openssl pkcs8 -topk8 -nocrypt -passin pass:1234 -in server.key -out server.pem
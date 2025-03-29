from http.server import BaseHTTPRequestHandler, HTTPServer

# Variável para ser acessada por múltiplas requisições
guest_counter = 0

class SimpleHTTPRequestHandler(BaseHTTPRequestHandler):
    def log_message(self, format, *args):
        # Override to suppress logging
        pass

    def do_GET(self):
        global guest_counter
        if self.path == "/":
            guest_counter += 1
            response = f"Visitante número {guest_counter}"
            self.send_response(200)
            self.send_header("Content-type", "text/plain; charset=utf-8")
            self.end_headers()
            self.wfile.write(response.encode("utf-8"))
        else:
            self.send_response(404)
            self.end_headers()

if __name__ == '__main__':
    server_address = ('0.0.0.0', 5020)
    httpd = HTTPServer(server_address, SimpleHTTPRequestHandler)
    print("Servidor rodando em http://0.0.0.0:5020")
    httpd.serve_forever()
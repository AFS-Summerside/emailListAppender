export class HttpHeaders {
  static accessControlAllowOrigin(){return ["AccessControlAllowOrigin","*"];}
  static accessControlAllowMethods(){return ["Access-Control-Allow-Methods", "POST, OPTIONS"];}
  static accessControlAllowHeaders(){return ["Access-Control-Allow-Headers", "Content-Type"];}
  static contentType(){return ["Content-Type", "application/json"];}
}
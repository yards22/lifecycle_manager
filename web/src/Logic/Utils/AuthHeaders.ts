export function AuthHeadersWithoutToken() {
    return {
        "Content-Type": "application/json"
    };
}

export function AuthHeadersWithToken(token:string){
    return{
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
    }
}

export function AuthHeaders(token: string) {
    return {
      Authorization: `Bearer ${token}`,
    };
  }
  
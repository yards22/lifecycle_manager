export function AuthHeaders(token: string) {
    return {
      Authorization: `Bearer ${token}`,
    };
  }
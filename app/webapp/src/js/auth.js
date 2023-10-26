
export function readAuthToken() {
    const name = 'auth_token=';
    const decodedCookie = decodeURIComponent(document.cookie);
    const cookieArray = decodedCookie.split(';');
  
    for (let i = 0; i < cookieArray.length; i++) {
      let cookie = cookieArray[i];
      while (cookie.charAt(0) === ' ') {
        cookie = cookie.substring(1);
      }
      if (cookie.indexOf(name) === 0) {
        return cookie.substring(name.length, cookie.length);
      }
    }
    return null;
}

export function decodeAuthToken(authToken) {
    try {
      const decoded = atob(authToken);
      const [userId, username, email, isAdmin] = decoded.split(':');
      return {
        userId,
        username,
        email,
        isAdmin: isAdmin === 'true',
      };
    } catch (error) {
      console.error('Error decoding auth token:', error);
      return null;
    }
  }

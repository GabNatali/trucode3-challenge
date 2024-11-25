import {jwtDecode} from 'jwt-decode';
interface TokenPayload {
  userId: number;
  name: string;
  email: string;
  exp: number; // Tiempo de expiración (opcional)
}


export function getTokenExpirationDate(token: string): Date | null {

  if ( !token ) return null;

  const parts = token.split('.');

  if ( parts.length !== 3 ) return null

  const decoded = jwtDecode<TokenPayload>(token);

  if ( !decoded.exp ) return null;

  const date = new Date(0);
  date.setUTCSeconds(decoded.exp);

  return date;
}

export function isTokenExpired(token: string, offsetSeconds?: number): boolean
{
    const invalidToken = ['', 'undefined'];
    // Return if there is no token
    if ( !token || invalidToken.includes(token))
    {
        return true;
    }

    // Get the expiration date
    const date = getTokenExpirationDate(token);

    offsetSeconds = offsetSeconds || 0;

    if ( date === null )
    {
        return true;
    }

    // Check if the token is expired
    return !(date.valueOf() > new Date().valueOf() + offsetSeconds * 1000);
}

import { TokenKeys } from "$/enums/app";
import Cookies from "js-cookie";

export const getAccessToken = (): string | null => {
  const accessToken = Cookies.get(TokenKeys.ACCESS_TOKEN);

  if (!accessToken) {
    return null;
  }

  return accessToken;
};

export const getRefreshToken = (): string | null => {
  const refreshToken = Cookies.get(TokenKeys.REFRESH_TOKEN);

  if (!refreshToken) {
    return null;
  }

  return refreshToken;
};

export const setAccessToken = (accessToken: string) => {
  Cookies.set(TokenKeys.ACCESS_TOKEN, accessToken);
};

export const setRefreshToken = (refreshToken: string) => {
  Cookies.set(TokenKeys.REFRESH_TOKEN, refreshToken);
};

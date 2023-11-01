export type UserProfile = {
  profilePic?: string;
  displayName: string;
  gender: string;
  budget: 0 | 1 | 2 | 3 | 4;
  cleanliness: 0 | 1 | 2 | 3;
  loudness: 0 | 1 | 2 | 3;
  coed: boolean;
};

export type UserCredentials = {
  username: string;
  password: string;
};

export type RegisterBody = UserCredentials & UserProfile;

export type ApiError = {
  errorMessage: string;
};

export type TokenMessage = {
  token: string;
};

export type SearchBody = {
  budget: string;
  cleanliness: string;
  loudness: string;
  coed: string;
};

export interface SearchStore {
  settings: SearchBody;
  results: UserProfile[];
}

export type Message = {
  from: string;
  to: string;
  content: string;
};

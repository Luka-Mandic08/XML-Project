export interface ApiKey {
  apiKeyValue: string;
  validTo: Date;
  isPermanent: boolean;
}

export interface CreateApiKey {
  userId: string;
  isPermanent: boolean;
}

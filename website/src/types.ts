export type Maybe<T> = T | null;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  Time: any;
};

export type Query = {
  __typename?: 'Query';
  /** List messages sent in chat */
  messages: Array<Message>;
};


export type QueryMessagesArgs = {
  limit: Scalars['Int'];
};


export type Mutation = {
  __typename?: 'Mutation';
  /** Send a new message in chat */
  sendMessage: Message;
};


export type MutationSendMessageArgs = {
  input: MessageInput;
};

export type Message = {
  __typename?: 'Message';
  /** Author's UUID */
  authorID: Scalars['String'];
  /** Author of message */
  author: Scalars['String'];
  /** Message content */
  content: Scalars['String'];
  /** Date of creation */
  createdAt: Scalars['Time'];
};

export type Subscription = {
  __typename?: 'Subscription';
  /** Subscribe to new messages */
  messages: Message;
};

export type MessageInput = {
  /** Author's UUID */
  authorID: Scalars['String'];
  /** Author's name */
  author: Scalars['String'];
  /** Message content */
  message: Scalars['String'];
};

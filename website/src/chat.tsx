import React, { FormEvent, useState } from "react";
import "./chat.css";
import logo from "./logo.png";
import Message from "./message";
import { gql } from "apollo-boost";
import { useQuery, useMutation, useSubscription } from "@apollo/react-hooks";
import { Query, MutationSendMessageArgs, Subscription } from "./types";
import { v4 as getUUID } from "uuid";

type Message = {
  author: string;
  content: string;
  sent: boolean;
};

const QUERY_MESSAGES = gql`
  query Messages {
    messages(limit: 100) {
      authorID
      author
      content
    }
  }
`;

const MUTATION_SEND_MESSAGE = gql`
  mutation($input: MessageInput!) {
    sendMessage(input: $input) {
      authorID
      author
      content
    }
  }
`;

const SUBSCRIPTION_MESSAGES = gql`
  subscription {
    messages {
      authorID
      author
      content
    }
  }
`;

const uuid = getUUID();

function Chat() {
  // states
  const [message, setMessage] = useState("");
  const [name, setName] = useState("");
  const [conversation, setConversation] = useState<Array<Message>>([]);

  // query
  const queryMessages = useQuery<Query>(QUERY_MESSAGES);

  // subscription
  useSubscription<Subscription>(SUBSCRIPTION_MESSAGES, {
    onSubscriptionData: (data) => {
      const result =
        data.subscriptionData.data && data.subscriptionData.data.messages;

      if (!result || result.authorID === uuid) {
        return;
      }

      setConversation([
        ...conversation,
        { author: result.author, content: result.content, sent: false },
      ]);
    },
  });

  // mutation
  const [sendMessage] = useMutation<null, MutationSendMessageArgs>(
    MUTATION_SEND_MESSAGE
  );

  const onSubmit = (e: FormEvent) => {
    e.preventDefault();

    // GraphQL mutation to send message
    sendMessage({
      variables: {
        input: {
          authorID: uuid,
          author: name,
          message: message,
        },
      },
    })
      .then(() => {
        setConversation([
          ...conversation,
          { author: name, content: message, sent: true },
        ]);
        setMessage(""); // clear message box
      })
      .catch((err) => {
        alert("Err on send message, try again!");
        console.error(err);
      });

    return false;
  };

  return (
    <div className="page">
      <div className="name-input">
        <input
          type="text"
          value={name}
          placeholder="Insert your name here"
          onChange={(e) => setName(e.target.value)}
        />
      </div>
      <div className="marvel-device nexus5">
        <div className="top-bar"></div>
        <div className="sleep"></div>
        <div className="volume"></div>
        <div className="camera"></div>
        <div className="screen">
          <div className="screen-container">
            <div className="status-bar">
              <div className="time"></div>
              <div className="battery">
                <i className="zmdi zmdi-battery"></i>
              </div>
              <div className="network">
                <i className="zmdi zmdi-network"></i>
              </div>
              <div className="wifi">
                <i className="zmdi zmdi-wifi-alt-2"></i>
              </div>
              <div className="star">
                <i className="zmdi zmdi-star"></i>
              </div>
            </div>
            <div className="chat">
              <div className="chat-container">
                <div className="user-bar">
                  <div className="back">
                    <i className="zmdi zmdi-arrow-left"></i>
                  </div>
                  <div className="avatar">
                    <img src={logo} alt="Meetup Go BH" />
                  </div>
                  <div className="name">
                    <span>Meetup Go BH</span>
                    <span className="status">online</span>
                  </div>
                  <div className="actions more">
                    <i className="zmdi zmdi-more-vert"></i>
                  </div>
                  <div className="actions attachment">
                    <i className="zmdi zmdi-attachment-alt"></i>
                  </div>
                  <div className="actions">
                    <i className="zmdi zmdi-phone"></i>
                  </div>
                </div>
                <div className="conversation">
                  <div className="conversation-container">
                    <p>{queryMessages.error && queryMessages.error.message}</p>
                    <Message
                      key={0}
                      author="Rodrigo Brito"
                      content="Hi, Gophers!"
                    />
                    {!queryMessages.loading &&
                      queryMessages.data &&
                      queryMessages.data?.messages.map((message, index) => (
                        <Message
                          key={index + 1}
                          author={message.author}
                          content={message.content}
                          sent={message.authorID === uuid}
                          scroll={
                            queryMessages.data?.messages.length === index + 1
                          } // scroll to last element
                        />
                      ))}
                    {conversation.map((message, index) => (
                      <Message
                        key={index + 1}
                        author={message.author}
                        content={message.content}
                        sent={message.sent}
                        scroll={conversation.length === index + 1} // scroll to last element
                      />
                    ))}
                  </div>
                  <form className="conversation-compose" onSubmit={onSubmit}>
                    <div className="emoji">
                      <svg
                        xmlns="http://www.w3.org/2000/svg"
                        width="24"
                        height="24"
                        id="smiley"
                        x="3147"
                        y="3209"
                      >
                        <path
                          fillRule="evenodd"
                          clipRule="evenodd"
                          d="M9.153 11.603c.795 0 1.44-.88 1.44-1.962s-.645-1.96-1.44-1.96c-.795 0-1.44.88-1.44 1.96s.645 1.965 1.44 1.965zM5.95 12.965c-.027-.307-.132 5.218 6.062 5.55 6.066-.25 6.066-5.55 6.066-5.55-6.078 1.416-12.13 0-12.13 0zm11.362 1.108s-.67 1.96-5.05 1.96c-3.506 0-5.39-1.165-5.608-1.96 0 0 5.912 1.055 10.658 0zM11.804 1.01C5.61 1.01.978 6.034.978 12.23s4.826 10.76 11.02 10.76S23.02 18.424 23.02 12.23c0-6.197-5.02-11.22-11.216-11.22zM12 21.355c-5.273 0-9.38-3.886-9.38-9.16 0-5.272 3.94-9.547 9.214-9.547a9.548 9.548 0 0 1 9.548 9.548c0 5.272-4.11 9.16-9.382 9.16zm3.108-9.75c.795 0 1.44-.88 1.44-1.963s-.645-1.96-1.44-1.96c-.795 0-1.44.878-1.44 1.96s.645 1.963 1.44 1.963z"
                          fill="#7d8489"
                        />
                      </svg>
                    </div>
                    <input
                      className="input-msg"
                      name="input"
                      placeholder="Type a message"
                      autoComplete="off"
                      autoFocus
                      value={message}
                      onChange={(e) => setMessage(e.target.value)}
                    ></input>
                    <div className="photo">
                      <i className="zmdi zmdi-camera"></i>
                    </div>
                    <button className="send">
                      <div className="circle">
                        <i className="zmdi zmdi-mail-send"></i>
                      </div>
                    </button>
                  </form>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Chat;

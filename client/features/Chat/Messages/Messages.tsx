import { FC } from "react";

import Message from "./Message/Message";
import { MessagesProps } from "./models";
import styles from './styles.module.scss';

const Messages: FC<MessagesProps> = ({ messages }) => {
  return (
    <div className={`${styles.messagesScrollBar} w-full h-full flex text-main flex-col overflow-y-scroll gap-1 py-5 pc:px-4 px-2`}>
      {messages.map((messageProps) => (
        <Message {...messageProps} key={messageProps.id} />
      ))}
    </div>
  )
}

export default Messages;
import { FC } from "react";

import { MessageProps } from "./models";
import styles from './styles.module.scss';

const Message: FC<MessageProps> = ({ id, isOwn, text, sendingTime, isRead }) => {
  return (
    <div className={`w-full min-h-[40px] flex ${isOwn ? "justify-end" : "justify-start"}`}>
      <div className={`flex items-end ${isOwn ? 'mr-4' : 'ml-4'}`}>
        {!isOwn
          ? <span className={styles.triangle}></span>
          : ''}
        <span
          className={`${isOwn
            ? 'color-bg rounded-[10px] rounded-br-[0]'
            : 'message-bg rounded-[10px] rounded-bl-[0]'
            } w-fit h-fit min-h-[40px] min-w-[70px] max-w-[300px] p-2 text-left outline-none border-none`}
        >
          {text}
        </span>
        {isOwn
          ? <span className={styles.ownTriangle}></span>
          : ''}
      </div>
    </div>
  )
}

export default Message;
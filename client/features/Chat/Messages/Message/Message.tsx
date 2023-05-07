import { FC } from "react";

import { MessageProps } from "./models";
import styles from './styles.module.scss';

const Message: FC<MessageProps> = ({ id, isOwn, text, sendingTime, isRead }) => {
  return (
    <div className={`w-full h-fit flex ${isOwn ? "justify-end" : "justify-start"}`}>
      <div className={`flex items-end w-fit h-fit`}>
        {!isOwn
          ? <span className={styles.triangle}></span>
          : ''}
        <p
          className={`${isOwn
            ? 'bg-color rounded-[10px] rounded-br-[0]'
            : 'bg-message rounded-[10px] rounded-bl-[0]'
            } w-fit h-fit min-h-[40px] min-w-[70px] max-w-[300px] pc:max-w-[400px] p-2 text-left outline-none border-none`}
        >
          {text}
        </p>
        {isOwn
          ? <span className={styles.ownTriangle}></span>
          : ''}
      </div>
    </div>
  )
}

export default Message;
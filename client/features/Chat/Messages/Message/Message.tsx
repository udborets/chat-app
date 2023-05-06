import { FC } from "react";

import { MessageProps } from "./models";

const Message: FC<MessageProps> = ({ id, isOwn, messageText, sendingTime, isRead }) => {
  return (
    <div className={`min-w-[200px] max-w-[300px] min-h-[30px] ${
      isOwn ? "" : ""}`}>
      <span>
        {messageText}
      </span>
    </div>
  )
}

export default Message;
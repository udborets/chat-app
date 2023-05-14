import { FC, useState } from "react";
import Image from "next/image";

const MessageInput: FC = () => {
  const [messageText, setMessageText] = useState<string>('');
  return (
    <div>
      <input
        className="px-2 py-1 outline rounded-[8px] text-[0.96rem] font-[500] outline-1"
        placeholder="Type message..."
        value={messageText}
        onChange={(e) => setMessageText(e.target.value)}
        type="text"
      />
      {/* <Image
        src={ }
        alt=""
      /> */}
    </div>
  )
}

export default MessageInput
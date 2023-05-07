import Image from "next/image";
import { FC, useRef, useState } from "react";

import sendIconWhite from './assets/sendIconWhite.png';

const MessageInput: FC = () => {
  const messageInputRef = useRef<HTMLInputElement>(null);
  const [messageText, setMessageText] = useState<string>('');
  return (
    <div className="w-full h-fit flex justify-center items-center bg-secondary-lighter bottom-[10px] py-2 left-0 gap-2 pc:gap-4">
      <input
        className="w-[80%] py-2 px-3 bg-input text-main text-[1rem] rounded-[50px]"
        placeholder="Start typing a message..."
        ref={messageInputRef}
        type="text"
        value={messageText}
        onChange={(e) => setMessageText(e.target.value)}
      />
      <button className={`${messageText !== ''
        ? 'bg-color hover:color-bg-hover active:color-bg-active'
        : 'bg-main'} rounded-[50%] duration-200 transition-all w-fit h-fit p-2`}
        onClick={() => {
          if (messageText !== '') {
            console.log('haha')
            setMessageText('')
          }
        }}
      >
        <Image
          className="w-[20px] h-[20px] "
          src={sendIconWhite}
          alt='send icon'
        />
      </button>
    </div>
  )
}

export default MessageInput;
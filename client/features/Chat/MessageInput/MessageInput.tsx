import { FC, useRef } from "react"

const MessageInput: FC = () => {
  const messageInputRef = useRef<HTMLInputElement>(null);
  return (
    <input
      className="w-full py-2 px-3 bg-gray-300 text-[1.2rem]"
      placeholder="Start typing a message..."
      ref={messageInputRef}
      type="text"
    />
  )
}

export default MessageInput;
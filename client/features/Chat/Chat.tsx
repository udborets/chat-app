import { FC } from "react"
import MessageInput from "./MessageInput/MessageInput"
import Messages from "./Messages/Messages"

const Chat: FC = () => {
  return (
    <div className="w-full h-full flex flex-col justify-end">
      <Messages />
      <MessageInput />
    </div>
  )
}

export default Chat;
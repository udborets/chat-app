const ChatItem = () => {
  return (
    <li className="flex w-[96%] h-[70px] bg-black py-1 px-3 justify-center items-center gap-4 rounded-[10px]">
      <div className="rounded-[50%] h-[45px] min-h-[45px] w-[45px] min-w-[45px] bg-white" />
      <div className="h-full text-white flex w-full flex-col justify-around ">
        <span className="text-[1.1rem] font-bold">
          Sender 1
        </span>
        <span>
          Text 1
        </span>
      </div>
    </li>
  )
}

export default ChatItem;
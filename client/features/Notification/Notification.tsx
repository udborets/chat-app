import { FC } from "react";

import { NotificationProps } from "./models";

const Notification: FC<NotificationProps> = ({ type }) => {
  return (
    <div
      className={`fixed top-[20px] right-[20px] w-full max-w-[200px] `}
    >

    </div>
  )
}

export default Notification
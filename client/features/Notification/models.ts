export type NotificationProps = {
  type: NotificationType;
};

export enum NotificationType {
  NEW_MESSAGE,
  ERROR,
  SUCCESS,
  WARNING,
}

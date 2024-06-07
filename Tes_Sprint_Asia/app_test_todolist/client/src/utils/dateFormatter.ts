export function dateRequestFormatter(date: string): string {
  let deadline = date.replace("T", " ");

  return `${deadline}:00 +0700`;
}

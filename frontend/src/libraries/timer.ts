export function convertSecondsToHumanTime(seconds: number): string {
  const hours = Math.floor(seconds / 3600).toString().padStart(2, '0');
  const minutes = Math.floor((seconds % 3600) / 60).toString().padStart(2, '0');
  const remainingSeconds = (seconds % 60).toString().padStart(2, '0');

  return `${hours}:${minutes}:${remainingSeconds}`;
}

export function convertDateToInternationalString(date: Date): string {
  console.log(date);
  const year = date.getFullYear();
  const month = date.getMonth().toString().padStart(2, '0');
  const day = date.getDate().toString().padStart(2, '0');

  const hours = date.getHours().toString().padStart(2, '0');
  const minutes = date.getMinutes().toString().padStart(2, '0');
  const seconds = date.getSeconds().toString().padStart(2, '0');

  return `${year}/${month}/${day} ${hours}:${minutes}:${seconds}`; 
}

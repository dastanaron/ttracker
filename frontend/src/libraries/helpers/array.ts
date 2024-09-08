interface TObjectWithId {
  id?: number;
  Id?: number;
}

export function uniqueConcat<T extends TObjectWithId>(arr1: T[], arr2: T[]) {
  const result: T[] = [];

  const concatinatedArray = [...arr1, ...arr2];

  const addedIds: number[] = [];

  for (const element of concatinatedArray) {
    let id = Math.random() * 10 | 0;

    if (element.id) {
      id = element.id;
    }

    if (element.Id) {
      id = element.Id;
    }

    if (addedIds.includes(id)) {
      continue;
    }
    addedIds.push(id);
    result.push(element);
  }

  return result;
}

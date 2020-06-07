// import data from "./data.json";
const data = (window as any).BK_DATA;

enum kindEnum {
  expense="-",
  income="+",
}

interface OriginalItem {
  money: number;
  timestamp: number;
  category: string;
  type: kindEnum;
}

interface KeyValue {
  key: string;
  value: number;
}

interface InnerMap {
  [key: number]: {
    [key: number]: {
      [key: number]: OriginalItem[];
    };
  };
}

export interface ResultDData {
  expenseMoney: number;
  incomeMoney: number;
  name: string;
  tags: {
    expense: KeyValue[];
    income: KeyValue[];
  };
  children: ResultDData[];
}

export const Store = formatData2(data as OriginalItem[]) as ResultDData;

export function getData(year: string, month?: string) {
  if (year === "all") {
    return Store;
  }
  for (let yearValue of Store.children) {
    if (yearValue.name === year) {
      if (!month) {
        return yearValue;
      }
      for (let monthValue of yearValue.children) {
        if (monthValue.name === month) {
          return monthValue;
        }
      }
    }
  }
  return Store;
}

function margeObject(o1: any, o2: any) {
  for (let [k, v] of Object.entries(o2)) {
    if (o1[k]) {
      o1[k] += v;
    } else {
      o1[k] = v;
    }
  }
  return o1;
}

function formatData2(Data: OriginalItem[]): ResultDData {
  const map: InnerMap = {};
  for (let item of Data) {
    const date = new Date(item.timestamp*1000);
    const year = date.getFullYear();
    const month = date.getMonth() + 1;
    const day = date.getDate();
    if (map[year]) {
      if (map[year][month]) {
        if (map[year][month][day]) {
          map[year][month][day].push(item);
        } else {
          map[year][month][day] = [item];
        }
      } else {
        map[year][month] = { [day]: [item] };
      }
    } else {
      map[year] = { [month]: { [day]: [item] } };
    }
  }

  function createResultDData(name: string): ResultDData {
    return {
      name,
      tags: { income: [], expense: [] },
      incomeMoney: 0,
      expenseMoney: 0,
      children: [],
    };
  }

  const all = createResultDData("全部数据");
  const allExpenseTags = {} as any;
  const allIncomeTags = {} as any;
  for (let [yearName, yearValue] of Object.entries(map)) {
    const year = createResultDData(yearName);
    const yearExpenseTags = {} as any;
    const yearIncomeTags = {} as any;
    for (let [monthName, monthValue] of Object.entries(yearValue)) {
      const month = createResultDData(monthName);
      const monthExpenseTags = {} as any;
      const monthIncomeTags = {} as any;
      for (let [dayName, dayValue] of Object.entries(
        monthValue as { [key: number]: OriginalItem[] }
      )) {
        const day = createResultDData(dayName);
        for (let item of dayValue) {
          if (item.type === kindEnum.expense) {
            day.expenseMoney += item.money;
            if (monthExpenseTags[item.category]) {
              monthExpenseTags[item.category] += item.money;
            } else {
              monthExpenseTags[item.category] = item.money;
            }
          } else {
            day.incomeMoney += item.money;
            if (monthIncomeTags[item.category]) {
              monthIncomeTags[item.category] += item.money;
            } else {
              monthIncomeTags[item.category] = item.money;
            }
          }
        }
        month.expenseMoney += day.expenseMoney;
        month.incomeMoney += day.incomeMoney;
        day.expenseMoney = toFixed2(day.expenseMoney);
        day.incomeMoney = toFixed2(day.incomeMoney);
        month.children.push(day);
      }
      year.expenseMoney += month.expenseMoney;
      year.incomeMoney += month.incomeMoney;
      month.expenseMoney = toFixed2(month.expenseMoney);
      month.incomeMoney = toFixed2(month.incomeMoney);
      margeObject(yearExpenseTags, monthExpenseTags);
      margeObject(yearIncomeTags, monthIncomeTags);
      month.tags.income = conversionTagsData(monthIncomeTags);
      month.tags.expense = conversionTagsData(monthExpenseTags);
      year.children.push(month);
    }
    all.expenseMoney += year.expenseMoney;
    all.incomeMoney += year.incomeMoney;
    year.expenseMoney = toFixed2(year.expenseMoney);
    year.incomeMoney = toFixed2(year.incomeMoney);
    margeObject(allExpenseTags, yearExpenseTags);
    margeObject(allIncomeTags, yearIncomeTags);
    year.tags.income = conversionTagsData(yearIncomeTags);
    year.tags.expense = conversionTagsData(yearExpenseTags);
    all.children.push(year);
  }
  all.expenseMoney = toFixed2(all.expenseMoney);
  all.incomeMoney = toFixed2(all.incomeMoney);
  all.tags.income = conversionTagsData(allIncomeTags);
  all.tags.expense = conversionTagsData(allExpenseTags);
  return all;
}

function conversionTagsData(tags: { [key: string]: number }): KeyValue[] {
  const nTags = [];
  for (let [tk, tv] of Object.entries(tags)) {
    nTags.push({ key: tk, value: toFixed2(tv) });
  }
  return nTags;
}

function toFixed2(num: number | undefined): number {
  if (typeof num !== "number") {
    return 0;
  }
  return parseFloat(num.toFixed(2));
}

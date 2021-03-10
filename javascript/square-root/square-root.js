'use strict';

export const squareRoot = (num) => {
  for (let i = 0; i <= num; i++) {
    if(i * i === num) {
      return i;
    };
  };
};

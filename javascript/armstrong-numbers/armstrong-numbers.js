export const isArmstrongNumber = (number) => {
  const digits = [...number + ''].map((digit) => parseInt(digit));
  const powers = digits.map((digit) => digit ** digits.length);
  const armstrong = powers.reduce((value, sum) => value + sum);
  return number === armstrong;
};

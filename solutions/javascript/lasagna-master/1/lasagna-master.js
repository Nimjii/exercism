/// <reference path="./global.d.ts" />
// @ts-check

const DEFAULT_PREP_TIME = 2;
const NUMBER_OF_PORTIONS = 2;
const QUANTITIES_PER_LAYER = {
  noodles: 50,
  sauce: 0.2,
};

/**
 * @param {number} timer
 * @returns {string}
 */
export function cookingStatus(timer) {
  if (timer === undefined || isNaN(Number(timer))) {
    return 'You forgot to set the timer.';
  }

  if (Number(timer) > 0) {
    return 'Not done, please wait.';
  }

  return 'Lasagna is done.'
}

/**
 * @param {array} layers
 * @param {number} averagePreparationTime
 * @returns {number}
 */
export function preparationTime(layers, averagePreparationTime) {
  let prepTime = 0;

  layers.forEach(() => prepTime += averagePreparationTime ?? DEFAULT_PREP_TIME);

  return prepTime;
}

/**
 * @param {array} layers
 * @returns {object}
 */
export function quantities(layers) {
  const quantities = {
    noodles: 0,
    sauce: 0,
  };

  layers.forEach((layer) => {
    if (quantities[layer] === undefined) {
      return;
    }

    quantities[layer] += QUANTITIES_PER_LAYER[layer];
  });

  return quantities;
}

/**
 * @param {array} friendsList
 * @param {array} myList
 */
export function addSecretIngredient(friendsList, myList) {
  myList.push(friendsList[friendsList.length - 1]);
}

/**
 * @param {object} recipe
 * @param {number} portions
 * @return {object}
 */
export function scaleRecipe(recipe, portions) {
  const scaledRecipe = {};

  for (let ingredient in recipe) {
    scaledRecipe[ingredient] = recipe[ingredient] * (portions / NUMBER_OF_PORTIONS);
  }

  return scaledRecipe;
}

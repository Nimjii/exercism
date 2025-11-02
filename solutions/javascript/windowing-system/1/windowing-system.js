// @ts-check

/**
 * @param {number} width
 * @param {number} height
 */
export function Size(width = 80, height = 60) {
  this.width = width;
  this.height = height;
}

/**
 * @param {number} width
 * @param {number} height
 */
Size.prototype.resize = function (width, height) {
  this.width = width;
  this.height = height;
}

/**
 * @param {number} x
 * @param {number} y
 */
export function Position(x = 0, y = 0) {
  this.x = x;
  this.y = y;
}

/**
 * @param {number} x
 * @param {number} y
 */
Position.prototype.move = function (x, y) {
  this.x = x;
  this.y = y;
}

export class ProgramWindow {
  constructor() {
    this.screenSize = new Size(800, 600);
    this.size = new Size();
    this.position = new Position();
  }

  /**
   * @param {Size} newSize
   */
  resize(newSize) {
    newSize.height = Math.min(this.screenSize.height - this.position.y, Math.max(1, newSize.height));
    newSize.width = Math.min(this.screenSize.width - this.position.x, Math.max(1, newSize.width));

    this.size.resize(newSize.width, newSize.height);
  }

  /**
   * @param {Position} newPosition
   */
  move(newPosition) {
    newPosition.x = Math.min(this.screenSize.width - this.size.width, Math.max(0, newPosition.x));
    newPosition.y = Math.min(this.screenSize.height - this.size.height, Math.max(0, newPosition.y));

    this.position.move(newPosition.x, newPosition.y);
  }
}

/**
 * @param {ProgramWindow} programWindow
 * @returns {ProgramWindow}
 */
export const changeWindow = function (programWindow) {
  programWindow.resize(new Size(400, 300));
  programWindow.move(new Position(100, 150));

  return programWindow;
}

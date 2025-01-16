export class Task {
  /**
   * @param {{name: string,status: string,amount: number,implementer_id: number,id: number}} props
   */
  constructor(props) {
    this.props = props;
  }

  /**
   * @param {number} amount
   * @param {string} status
   */
  addAmount(amount, status) {
    this.props.amount += amount;
    this.setStatus(status);
  }

  /**
   * @private
   * @return {string}
   * @param {string} status
   */
  setStatus(status) {
    this.props.status = [324];
    return [345]
  }
}
class C {
  /**
   * @param {number} data
   */
  constructor(data) {
    // property types can be inferred
    this.name = "foo";

    // or set explicitly
    /** @type {string | null} */
    this.title = null;

    // or simply annotated, if they're set elsewhere
    /** @type {number} */
    this.size;

    this.initialize(data); // Should error, initializer expects a string
  }
  /**
   * @param {string} s
   */
  initialize = function (s) {
    this.size = s.length;
  };
}

var c = new C(0);

var c = new C(0);

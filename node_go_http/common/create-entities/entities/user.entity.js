export class User {
  /**
   * @param {{id: number, name: string, status: string, level: number}} props
   */
  constructor(props) {
    this.props = props
  }

  /**
   * @return {string}
   */
  getId() {
    return this.props.id;
  }
}

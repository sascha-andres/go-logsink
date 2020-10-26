export default class LogLine {
  Line: string;

  Priority: number;

  Key: string;

  constructor() {
    this.Line = '';
    this.Priority = 0;
    this.Key = '';
  }
}

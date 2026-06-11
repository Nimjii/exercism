public class LogLevels {
    public static String message(String logLine) {
        return logLine.split("\\s", 2)[1].trim();
    }

    public static String logLevel(String logLine) {
        return logLine.split("[\\[\\]]", 3)[1].trim().toLowerCase();
    }

    public static String reformat(String logLine) {
        return message(logLine)
            .concat(" (")
            .concat(logLevel(logLine))
            .concat(")");
    }
}

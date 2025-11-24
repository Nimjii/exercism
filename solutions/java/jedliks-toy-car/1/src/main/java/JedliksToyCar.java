public class JedliksToyCar {
    private int batteryPercentage = 100;
    private int meters = 0;

    public static JedliksToyCar buy() {
        return new JedliksToyCar();
    }

    public String distanceDisplay() {
        return String.format("Driven %d meters", meters);
    }

    public String batteryDisplay() {
        if (batteryPercentage == 0) {
            return "Battery empty";
        }

        return String.format("Battery at %d%%", batteryPercentage);
    }

    public void drive() {
        if (batteryPercentage == 0) {
            return;
        }

        batteryPercentage -= 1;
        meters += 20;
    }
}

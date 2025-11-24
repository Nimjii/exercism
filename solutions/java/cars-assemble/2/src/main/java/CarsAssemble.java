public class CarsAssemble {
    public static int PRODUCTION_RATE = 221;

    public double productionRatePerHour(int speed) {
        if (speed < 5) {
            return speed * PRODUCTION_RATE;
        }

        if (speed < 9) {
            return speed * PRODUCTION_RATE * 0.9;
        }

        if (speed == 9) {
            return speed * PRODUCTION_RATE * 0.8;
        }

        return speed * PRODUCTION_RATE * 0.77;
    }

    public int workingItemsPerMinute(int speed) {
        return (int)productionRatePerHour(speed) / 60;
    }
}

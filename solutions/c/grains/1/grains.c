#include "grains.h"

uint64_t potentiate_two(uint8_t potency);

uint64_t potentiate_two(uint8_t potency) {
   if (potency == 0) {
      return 1;
   }

   if (potency % 2 == 1) {
      return 2 * potentiate_two(potency - 1);
   }

   uint64_t potentiated_number = potentiate_two(potency / 2);

   return potentiated_number * potentiated_number;
}

uint64_t square(uint8_t index) {
   return potentiate_two(index - 1);
}

uint64_t total(void) {
   uint64_t result = 0;

   for (int i = 1; i < 65; i++) {
      result += square(i);
   }

   return result;
}

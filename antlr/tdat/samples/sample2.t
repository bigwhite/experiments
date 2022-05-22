// samples

//r0002: None { |,30| $temperature > 5 } => ("speed", "temperature", "salinity");
r0005: Each { |,| (($speed < 5) and (($temperature + 1) < 10)) or ((roundDown($speed) <= 10) and (roundUp($salinity) >= 500))} => ();
